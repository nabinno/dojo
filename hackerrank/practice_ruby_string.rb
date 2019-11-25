##
# Encoding
def transcode(str)
  str.force_encoding(Encoding::UTF_8)
end

##
# Indexing
def serial_average(str)
  values = str.split('-')
  average = ((values[1].to_f + values[2].to_f) / 2).round(2)
  "#{values[0]}-#{average}"
end

##
# Iteration
def count_multibyte_char(ss)
  rc = 0
  ss.split("")
    .each do |char|
    rc += 1 if char.bytesize > 1
  end
  rc
end

##
# Methods I
def process_text(args)
  args.map(&:strip).join(" ")
end

##
# Methods II
def strike(str)
  "<strike>#{str}</strike>"
end

def mask_article(str, strikes)
  strikes.each do |strike|
    str = str.gsub(/(#{strike})/, "<strike>\\1</strike>")
  end
  str
end
