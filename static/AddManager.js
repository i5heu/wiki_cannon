var AddManagerHtml = `
<b>AddManager</b><br><br>
<button id="AddShortcut">Add Shortcut</button>
`

var AddShortcutHtml = `
<b>AddShortcut</b><br><br>
<input id="AddManagerTitle1">Title</input><br>
<input id="AddManagerText1">Link</input><br>
<input id="AddManagerNum1" type="number" >Priority</input><br>
<button id="AddShortcutSend">Send</button>
`


$( document ).ready(function() {

  $("#AddManager").mousedown(function() {
    defaultModalShow();
    $("#DefaultModalContent").html(AddManagerHtml)
  });

  $('#DefaultModalContainer').on( 'click', '#AddShortcut', function () {
    $("#DefaultModalContent").html(AddShortcutHtml)
  });

  $('#DefaultModalContainer').on( 'click', '#AddShortcutSend', function () {
    AddManagerSend();
  });

}); // DO NOT REMOVE DOC RDY


function AddManagerSend(){


  var xhr = new XMLHttpRequest();
  var url = "/api2";
  xhr.open("POST", url, true);
  xhr.setRequestHeader("Content-type", "application/json");
  xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200) {
          var json = JSON.parse(xhr.responseText);
          console.log(json.Status);
          $("#DefaultModalContent").html(json.Status);

          setTimeout(function(){
            defaultModalHide();
          }, 800);

      }
  };

  var data = JSON.stringify({ "APP":"ItemWrite", "PWD":$.cookie("pwd"), "APPWRITE":"shortcut", "Title1": $("#AddManagerTitle1").val(), "Text1": $("#AddManagerText1").val(),"Num1": parseInt($("#AddManagerNum1").val(),10 ) });
  console.log(data);

  xhr.send(data);
    $("#DefaultModalContent").html(`Sending <div class="LodingIndicator"></div>`)

}
