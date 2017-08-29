//ProjectManager

function GetProject(ProjectID){

    var xhr = new XMLHttpRequest();
    var url = "/api2";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var json = JSON.parse(xhr.responseText);
            console.log(json.DATA);
            $("#DefaultModalContent").html(`<table class="fancytable" id="ProjectReaderModalcontent"><tr><th>ItemID</th><th>Timecreate</th><th>Title1</th><th>Title2</th><th>Text1</th><th>Text2</th><th>Tags1</th><th>Num1</th><th>Num2</th><th>NNum3</th><th>Finsh</th><th>DEL</th></tr></table>`)

            $(json.DATA).each(function(index, item) {
                var t = `<tr><td>` + item.ItemID + `</td><td>` + item.Timecreate + `</td><td>` + item.Title1 + `</td><td>` + item.Title2 + `</td><td>` + item.Text1 + `</td><td>` + item.Text2 + `</td><td>` + item.Tags1 + `</td><td>`  + item.Num1 + `</td><td>` + item.Num2 + `</td><td>` + item.Num3 + `</td><td><button class="ProjectChildButton" data-methode="FINISH" data-id="` + item.ItemID + `">✔</button></td><td><button class="ProjectChildButton" data-methode="DEL" data-id="` + item.ItemID + `">X</button></td></tr>`

                $("#ProjectReaderModalcontent").append(t);
            });
        }
    };

    var data = JSON.stringify({ "APP":"ProjectRead", "PWD":$.cookie("pwd"), "ID": parseInt(ProjectID,10 ) });
    console.log(data);

    xhr.send(data);
      $("#DefaultModalContent").html(`Loading <div class="LodingIndicator"></div>`)

}


function AjaxProjectTask(APP,data,id){

    $("#LoadingIndicator").show();
    defaultModalShow();

    var xhr = new XMLHttpRequest();
    var url = "/api2";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var json = JSON.parse(xhr.responseText);
            $("#LoadingIndicator").hide();
            console.log(json.Status);
            GetProject(id);
            return true
            }
      };

        console.log(data);

        xhr.send(data);


  }

$( document ).ready(function() {
  $( ".ProjectManagerModalInitializer" ).click(function() {
     var foo = $(this).data("projectid")
     defaultModalShow();
     GetProject(foo);
     $('#DefaultModalContent').data('id',foo);
  });

  $('#DefaultModalContent').on( 'click', '.ProjectChildButton', function () {
    var TaskID = $(this).data("id")
    var Methode = $(this).data("methode")

    if (Methode == "DEL") {
      var data = JSON.stringify({ "APP":"ItemDelete", "PWD":$.cookie("pwd"),"ID":parseInt(TaskID,10)})
    }

    if (Methode == "FINISH") {
      alert("UNDER DEVOLOPMENT")
    }

    AjaxProjectTask("ProjectTaskEdit",data,$('#DefaultModalContent').data("id") );


  });


}); // DO NOT REMOVE DOC RDY
