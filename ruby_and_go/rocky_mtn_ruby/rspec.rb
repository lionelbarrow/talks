require "spec_helper"

describe MyClass do
  describe "this_great_method" do
    it "does some stuff" do
      m = MyClass.New

      result = m.this_great_method

      result.should do_some_stuff
    end
  end
end
