require 'sinatra'
require 'sequel'

# connect to an in-memory database
DB = Sequel.sqlite
items = DB[:items]

get "/:echo" do
  "#{params[:echo]}"
  items.insert(:name => 'abc', :price => rand * 100)
end

