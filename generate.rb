require 'cast'
require 'pry'

def camelize name
  parts = name.split("_").map(&:downcase).map do |part|
    part[0] = part[0].upcase
    part
  end
  return parts.join("")
end

lines = File.read('/usr/include/ta-lib/ta_func.h').split("\n")

recent_comment = ""
recent_func = ""
in_comment = true
in_func = false

funcs = []

$types = {
  "double"=> "float64"
}

class Func
  def initialize comment, func
    #@comment_str = comment
    #@func_str = func
    @name_raw = func.match(/ TA_(\w+)\(/)[1]
    @name = camelize @name_raw
    @args = func.match(/\(.*\)/)[0][1..-2].strip.split(",").map{|a| a.split(" ")}
    @comment = "/*#{@name} - " +comment.gsub(/\/\*(\n\*\s*\w*( - )?)?/, '').gsub(/^\s*\*\s+/,"").gsub(/ +/, " ").gsub(@name_raw, @name).strip
  end
  def to_go
    args = []
    body = []
    params = []
    returns = []
    returnTypes = []
    i = 0
    @args.each do |arg_set|
      type = arg_set[-2] #arg_set[0..-2].join(" ")
      arg = arg_set[-1]
      if arg == "startIdx"
        params << "0"
      elsif arg == "endIdx"
        params << "C.int(len(#{@args[i+1][1]}-1))"
      elsif arg.start_with? "*out"
        body << "var #{arg[1..-1]} C.#{type}"
        params << "&#{arg[1..-1]}"
      elsif arg.start_with?("in") || arg.start_with?("optIn")
        param = arg[2..-1].match(/(\w*)(\[\])?/)[1]
        if arg.start_with? "optIn"
          param = arg[5..-1].match(/(\w*)(\[\])?/)[1]
        end
        param[0] = param[0].downcase
        param.gsub!("_", "")
        goType = $types[type.split(" ").last]
        goType = type if !goType
        fType = ""
        if arg.end_with? "[]"
          fType += "[]"
          params << "(*C.#{type})(unsafe.Pointer(&#{param}[0]))"
        else
          params << param
        end
        fType += goType
        if args.last && args.last.end_with?(" "+fType)
          args[-1] = args.last[0...-(fType.length + 1)] + ", "+param + " " + fType
        else
          args << param + " " + fType
        end
      elsif arg.start_with? "out"
        param = arg[3..-1].match(/(\w*)(\[\])?/)[1]
        param[0] = param[0].downcase
        goType = $types[type.split(" ").last]
        goType = type if !goType
        if arg.end_with? "[]"
          body << "#{param} := make([]#{goType}, len(#{args.first.split(" ").first}))"
          params << "(*C.#{type})(unsafe.Pointer(&#{param}[0]))"
          returns << param
          returnTypes << "[]#{goType}"
        else
          params << "(*C.#{type})(unsafe.Pointer(&#{param}))"
          returns << param
          returnTypes << goType
        end
      end

      i += 1
    end
    s = @comment + "\n"
    s += "func #{@name}"
    s += "(" + args.join(", ") + ")"
    if returnTypes.length > 0
      s += " (#{returnTypes.join(", ")})"
    end
    s += " {\n"
    s += body.join("\n")
    s += "\n"
    s += "C.#{@name_raw}(#{params.join(", ")})"
    s += "\nreturn #{returns.join(", ")}\n"
    s += "}\n"
    s
  end
end

lines.each do |line|
  line = line.strip
  if line.start_with? "/*"
    recent_comment = ""
    in_comment = true
  end
  if in_comment
    recent_comment += line.strip + "\n"
  end
  if line.include? "*/"
    in_comment = false
  end
  if line.start_with? "TA_RetCode"
    recent_func = ""
    in_func = true
  end
  if in_func
    recent_func += line + " "
  end
  if line.include? ");"
    in_func = false
    if (!recent_func.start_with?("TA_RetCode TA_S_"))&& recent_func.start_with?("TA_RetCode TA_")
      funcs.push(Func.new(recent_comment, recent_func))
    end
  end
end


code = "package talib
// #cgo LDFLAGS: -lta_lib
// #include \"ta-lib/ta_libc.h\"
import \"C\"

func init() {
	n, err := C.TA_Initialize()
	if n != 0 {
		panic(fmt.Sprintf(\"ta-lib status is %d %s\", n, err))
	}
}
"
code += funcs.map(&:to_go).join("\n")

File.write("generated.go", code)
system("go fmt ./generated.go")
