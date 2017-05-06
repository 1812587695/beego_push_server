# beego_push_server
beego gorilla websocket消息推送功能

# js接受
  <script>
        var socket = new WebSocket("ws://192.168.210.228:8080/ws/join?store_id=12");

        socket.onopen = function (event) {
            console.log("Socket opened successfully");
        };
        socket.onmessage = function (event) {

            var json = eval('(' + event.data + ')');
            var j = eval('(' + json + ')');
            if (j.data != null){
                var num = j.data.length;
                $(".order_msg").text(num);

                var html = '';
                for (var i = 0; i < num; i++) {
                    html += '<li id="websoket_'+ j.data[i].id +'"><a href="#" onclick="deleteOrder(' + j.data[i].id + ')" style="float: right;">删除</a>';
                    html += '<a href="#"><i class="fa fa-shopping-cart text-green"></i> ' + j.data[i].msg + '</a></li>';
                }
                $('.order-ul-msg').append(html);
            }
        };

        window.onbeforeunload = function (event) {
            socket.close();
        };
     

    </script>
