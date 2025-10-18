require 'bundler/setup'
require 'sinatra'
require 'json'
require_relative 'lib/user_interface'
require_relative 'lib/playfair_operations'

# ui = UserInterface.new

set :port, 4567
set :bind, '0.0.0.0'

get '/' do
  send_file 'client/index.html'
end

post '/encrypt' do
  content_type :json

  data = JSON.parse(request.body.read)
  key = data['key']
  text = data['text']

  return { error: 'Key and text are required' }.to_json if key.nil? || text.nil? || key.empty? || text.empty?

  playfair = PlayfairOperations.new
  result = playfair.encrypt(key, text)

  { result: result }.to_json
end

post '/decrypt' do
  content_type :json

  data = JSON.parse(request.body.read)
  key = data['key']
  ciphertext = data['ciphertext']

  return { error: 'Key and ciphertext are required' }.to_json if key.nil? || ciphertext.nil? || key.empty? || ciphertext.empty?

  playfair = PlayfairOperations.new
  result = playfair.decrypt(key, ciphertext)

  { result: result }.to_json
end