//################## PwdManager  ###################



$( document ).ready(function() {

  $("#PwdManagerTrigger").mousedown(function() {
    PwdManager();
  });

  $("#DefaultModalClose").mousedown(function() {
    defaultModalHide();
  });

  $('#DefaultModalContainer').on( 'click', '.AddPassword', function () {
    PwdManagerAddview();
  });



}); // DO NOT REMOVE DOC RDY


function PwdManager(){
  defaultModalShow();
  data = '{"PWD":"'+  $.cookie("pwd") + `", "APP":"PwdManager"}`;

  $.ajax({
              type:"POST",
              url: "/api2",
              data:data,
              success: function (response){
                    var json = $.parseJSON(response);
                    console.log("PwdManager");
                    $("#DefaultModalContent").html(`<button class="AddPassword" >Add Password</button><br><br><table>`)
                    $(json.PwdResult).each(function(index, item) {

                        foo = "<tr><th>"+item.title1+"</th><th>"+item.title2+"</th><th>"+item.text1+"</th></tr>"
                        $("#DefaultModalContent").append(foo)
                    });
                    $("#DefaultModalContent").append(`</table>`)
                  }
        });
}


function PwdManagerAddview(){
  defaultModalShow();
  $("#DefaultModalContent").html(PwdInput)

}

function PwdManagerAddAPI(){

  data = '{"PWD":"'+  $.cookie("pwd") + `", "APP":"PwdManager"}`;

  $.ajax({
              type:"POST",
              url: "/api2",
              data:data,
              success: function (response){
                    var json = $.parseJSON(response);
                    console.log("PwdManager");
                    $("#DefaultModalContent").html(json.status)
                    setTimeout( function() { PwdManager() }, 1000);
                  }
        });
}





//////////////////// HTML /////////////////////////////
var PwdInput = `
<form>
<ul class="PwdInput">
<li class="PwdInput">
  <label class="PwdInput" for="text_id">Site</label>
  <input type="text" class="PwdInput" name="site" id="PwdInput-site" value="" />
</li>
<li class="PwdInput">
  <label class="PwdInput" for="text_id">username</label>
  <input type="text" class="PwdInput" name="username" id="PwdInput-username" value="" />
</li>
<li class="PwdInput">
  <label class="PwdInput" for="text_id">passsword</label>
  <input type="text" class="PwdInput" name="passsword" id="PwdInput-password" value="" />
</li>
</ul>
</form>`;
