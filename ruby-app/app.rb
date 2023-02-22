
$file_path = "../super-important-data.txt"

class App
    def app()
        data = File.read($file_path)
        return data + ", but in ruby."
    end
end