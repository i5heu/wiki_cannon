$( document ).ready(function() {

  //FOCUS ON KEYPRESS
  $(document).keyup(function(e) {
  if (! ($(e.target).is('input, textarea, .ed, #editarea, .divinput, .textarea, #tags__, #namespace_, #slug_, .texta'))) { //NOT TRIGGER IF INPUT ETC.

 // console.log(e.keyCode);

    switch(e.keyCode){
      case 32:
          openGO();
          $("#goto").focus();
        break;
      case 80:
          PwdManagerMaterPW();
        break;
      case 27:
          defaultModalHide();
        break;
    }


  }
  });



}); // DO NOT REMOVE DOC RDY


function openGO(){
  $( "#go_button" ).hide();
  $(".goto").css("visibility", "visible");
  $(".goto").css("display", "flex");
  $(".goto").css("opacity", "1");
  $("#clickscroll").css("display", "none");
  $(".close_button").css("display", "block");
  $("#goto").focus();
  if (screen.width < 550) {
    $("#afterGO").hide();
  }
}
