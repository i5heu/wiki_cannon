//################## PwdManager  ###################
var pwdhash

$( document ).ready(function() {

  $("#PwdManagerTrigger").mousedown(function() {
    PwdManagerMaterPW();
  });

  $('#DefaultModalContainer').on( 'click', '#PwdManagerTriggerAfterMaterPW', function () {
    PwdManagerMaterPWSave();
    PwdManager();
  });

  $("#DefaultModalClose").mousedown(function() {
    pwdhash = 0
    defaultModalHide();
  });

  $('#DefaultModalContainer').on( 'click', '.AddPassword', function () {
    PwdManagerAddview();
  });

  $('#DefaultModalContainer').on( 'click', '.SubmitPassword', function () {
    PwdManagerAddAPI($('#PwdInput-site').val(),$('#PwdInput-username').val(),$('#PwdInput-password').val(),);
  });



}); // DO NOT REMOVE DOC RDY



function PwdManagerMaterPW(){
  defaultModalShow();
  $("#DefaultModalContent").html(`<input id="masterpw">MasterPWD</input> <button id="PwdManagerTriggerAfterMaterPW">SHOW ME</button>`)
}

function PwdManagerMaterPWSave(){
  var foo = $('#masterpw').val();
  pwdhash = CryptoJS.SHA256(foo).toString(CryptoJS.enc.Hex);

}



function PwdManager(){
  defaultModalShow();
  data = '{"PWD":"'+  $.cookie("pwd") + `", "APP":"PwdManager"}`;

  $.ajax({
              type:"POST",
              url: "/api2",
              data:data,
              success: function (response){
                    var json = $.parseJSON(response);
                    $("#DefaultModalContent").html(` `)
                    $("#DefaultModalContent").append(`<button class="AddPassword" >Add Password</button><br><br> <table class="fancytable"><tr><td>Site</td><td>Username</td><td>Password</td></tr></table>`)
                    var switcher = "tablelight"

                    $(json.PwdResult).each(function(index, item) {
                        var decrypted = CryptoJS.AES.decrypt(item.text1, pwdhash);

                        foo = "<tr class='" + switcher + "'><th>"+item.title1+"</th><th>"+item.title2+"</th><th>"+decrypted.toString(CryptoJS.enc.Utf8)+"</th></tr>"


                        $('#DefaultModalContent tr:last').after(foo);


                        if(switcher == "tablelight") {
                          switcher = "tabledark"
                        }else{
                          switcher = "tablelight"
                        }
                    });

                  }
        });
}



function PwdManagerAddview(){
  defaultModalShow();
  $("#DefaultModalContent").html(PwdInput)

}

function PwdManagerAddAPI(site,username,passsword){


  var encrypted = CryptoJS.AES.encrypt(passsword, pwdhash);

  data = '{"PWD":"'+  $.cookie("pwd") + `", "APP":"ItemWrite"` + `, "APPWRITE":"PwdManager","Title1":"` + site + `","Title2":"` + username + `"` + `,"Text1":"`+encrypted+`"` +`}`;

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
<button class="SubmitPassword" >SubmitPassword</button>
`;
