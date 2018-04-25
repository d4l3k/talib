require 'cast'
require 'pry'

$toCapitalize = %w(
Price Bands Crows Black Inside Line Strike Stars In South Outside Conceal
White Soldiers Abandoned Baby Advance Block Belthold Breakaway Closing Marubozu
Swall Counterattack Dark Cloud Cover Doji Star Dragonfly Engulfing Evening
Gap Side Gravestone Hammer Hanging Man Harami Cross High Wave Hikkake Mod
Homing Pigeon Identical Neck Kicking By Length Ladder Bottom Long Legged
Matching Low Mat Hold Morning On Piercing Rickshaw Rise Fall Methods Separating
Shooting Short Spinning Top Staled Pattern Stick Sandwich Tasuki Takuri Stalled
Thrusting Tristar Unique River Upside Period Phase Sine Reg Linear Fix Ext Point
Min Index Max Sin Dev Rsi TriMa Osc Asin
).sort_by(&:size)

def camelize name
  parts = name.downcase.split("_").map do |dp|
    part = dp.dup
    $toCapitalize.each do |cap|
      i = dp.index(cap.downcase)
      if i
        part[i...(i+cap.length)] = cap
      end
    end
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
  "double" => "float64",
  "TA_MAType" => "int32",
  "int" => "int32",
}

class Func
  def initialize comment, func
    #@comment_str = comment
    #@func_str = func
    @name_raw = func.match(/ TA_(\w+)\(/)[1]
    @name = camelize @name_raw
    @args = func.match(/\(.*\)/)[0][1..-2].strip.split(",").map{|a| a.split(" ")}
    @comment = "/*#{@name} - " +comment.gsub(/\/\*(\n\*\s*\w*( - )?)?/, '').gsub(/^\s*\*\s+/,"").gsub(/ +/, " ").gsub(@name_raw, @name).strip.gsub("\n", "\n\n")
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
        param = @args[2].last[2..-1].match(/(\w*)(\[\])?/)[1]
        param[0] = param[0].downcase
        params << "C.int(len(#{param})-1)"
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
          params << "C.#{type}(#{param})"
        end
        fType += goType
        if args.last && args.last.end_with?(" "+fType)
          args[-1] = args.last[0...-(fType.length + 1)] + ", "+param + " " + fType
        else
          args << param + " " + fType
        end
      elsif arg.start_with? "out"
        param = arg.match(/(\w*)(\[\])?/)[1]
        param[0] = param[0].downcase
        goType = $types[type.split(" ").last]
        goType = type if !goType
        if arg.end_with? "[]"
          body << "#{param} := make([]#{goType}, len(#{args.first.split(" ").first}))"
          params << "(*C.#{type})(unsafe.Pointer(&#{param}[0]))"
          returns << "#{param}[:outNBElement]"
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
    s += "C.TA_#{@name_raw}(#{params.join(", ")})"
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
    func_name = recent_func.match(/TA_RetCode (\w+)\(/)[1]
    if (!recent_func.start_with?("TA_RetCode TA_S_")) && func_name == func_name.upcase
      funcs.push(Func.new(recent_comment, recent_func))
    end
  end
end


code = "package talib
// #cgo LDFLAGS: -lta_lib
// #include \"ta-lib/ta_libc.h\"
import \"C\"

import (
  \"fmt\"
  \"unsafe\"
)

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
