require_relative "app"
require "test/unit"
 
class TestApp < Test::Unit::TestCase

  def test_simple
    actual = App.new.app()
    assert_equal("Man, I really love cats, but in ruby.", actual)
  end

end