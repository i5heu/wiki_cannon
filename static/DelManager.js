
$( document ).ready(function() {


  $("#DelManager").mousedown(function() {
    defaultModalShow();
    LoadAllItems();
    ToogleMenue();
  });

  $('#DefaultModalContent').on( 'click', '.DelManagerItemDelButton', function () {
    var TaskID = $(this).data("id");
    var Methode = $(this).data("methode");

    if (Methode == "DEL") {
      var data = JSON.stringify({ "APP":"ItemDelete", "PWD":$.cookie("pwd"),"ID":parseInt(TaskID,10)});
    }


     ItemAllSend(data);



  });


});


function LoadAllItems(){
  $("#LoadingIndicator").show();
  data = '{"PWD":"'+  $.cookie("pwd") + `", "APP":"ItemAll"}`;

  $.ajax({
              type:"POST",
              url: "/api2",
              data:data,
              success: function (response){
                    var json = $.parseJSON(response);
                    $("#DefaultModalContent").html(` `)
                    $("#DefaultModalContent").append(`<table class="fancytable"><tr><th>Site</th><th>Username</th><th>Password</th></tr></table>`)

                    var switcher = "tablelight"

                    $(json.DATA).each(function(index, item) {


                        foo = "<tr class='" + switcher + "'><td>"+item.ItemID+"</td><td>"+item.APP+"</td><td>"+item.Title1+"</td><td>"+item.Title2+"</td><td>"+item.Text1+"</td><td>"+item.Text2+"</td><td>"+item.Tags1+"</td><td>"+item.Num1+"</td><td>"+item.Num2+"</td><td>"+item.Num3+"</td><td>"+item.Timecreate+"</td><td><button class='DelManagerItemDelButton' data-methode='DEL' data-id='"+item.ItemID+"'>X</button></td></tr>"


                        $('#DefaultModalContent tr:last').after(foo);


                        if(switcher == "tablelight") {
                          switcher = "tabledark"
                        }else{
                          switcher = "tablelight"
                        }

                    });
                    $("#LoadingIndicator").hide();

                  }
        });
}


function ItemAllSend(data){

    $("#LoadingIndicator").show();

    var xhr = new XMLHttpRequest();
    var url = "/api2";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var json = JSON.parse(xhr.responseText);
            $("#LoadingIndicator").hide();
            console.log(json.Status);

            LoadAllItems();
            return true
            }
      };

        console.log(data);

        xhr.send(data);

  }
