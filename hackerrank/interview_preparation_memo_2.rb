require 'faraday'
require 'faraday_middleware'

def askServer(n)
  path = "/api/recursive/ask?seed=#{@seed}&n=#{n}"
  res = @conn.get { |r| r.url(URI.encode(path)) }
  res.body['result']
end

def genNs(n)
  case n
  when 0 then @ns[n] += 1
  when 2 then @ns[n] += 1
  else
    if n % 2 == 0
      genNs(n - 1) + genNs(n - 2) + genNs(n - 3) + genNs(n - 4)
    else
      @ns[n] += 1
    end
  end
end

def newF()
  @ns.map(&lambda do |k,v|
            case k
            when 0 then 1 * v
            when 2 then 2 * v
            else
              askServer(k) * v
            end
          end)
    .sum
end

# @deprecated
def f(n)
  case n
  when 0 then 1
  when 2 then 2
  else
    if n % 2 == 0
      f(n - 1) + f(n - 2) + f(n - 3) + f(n - 4)
    else
      askServer(n)
    end
  end
end

def main(argv)
  if argv.nil? || argv.size == 0
    raise 'argv nil or argv.size == 0'
  elsif argv.size == 1
    raise 'argv.size == 1'
  elsif argv.size > 2
    raise 'argv.size > 2'
  elsif !argv[1].is_a?(Integer) && argv[1].to_i == 0 && argv[1] != '0'
    raise 'argv[1] not integer'
  end

  @seed = argv[0]
  n = argv[1].to_i
  endpoint = "http://challenge-server.code-check.io"
  @conn = Faraday.new(url: endpoint) do |f|
    f.request :json
    f.response :json, :content_type => /\bjson$/
    f.adapter :net_http
  end

  @ns = Hash.new{0}
  genNs(n)

  puts newF()
end

begin
  main(ARGV)
  # main(['xxx', '0'])
rescue => e
  puts "error! #{e}"
  exit 1
end

