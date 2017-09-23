def proto_task(name, &block)
    task name do
      orig_path = ENV['PATH']
      ENV['PATH'] = "#{Dir.pwd}/bin:#{Dir.pwd}/go/bin:#{ENV['PATH']}"
      Dir.chdir("go/src") do
        block.call
      end
      ENV['PATH'] = orig_path
    end
    task :all => name
  end

proto_task :location do
    sh "protoc --gogo_out=. -I=. app/proto/location/location.proto"
end
