$( document ).ready(function() {

  //CALIBRATING SPACER FOR HEADER
  function resizeHeader(){
   var header = $("header").height();
   document.getElementById("spacer").style.height = header + "px";
  }

  //CALIBRATING SPACER FOR HEADER on pageload
  window.onload = function () {
   resizeHeader();
  }


  $( ".fancybutton" ).click(function() {//OPEN
    $( "#HotkeyModalContainer" ).toggle();
  });

  $( "#HotkeyModalClose" ).click(function() {//OPEN
    $( "#HotkeyModalContainer" ).hide();
  });




  //MOBILE MENUE BUTTON
  $( ".menue-button" ).click(function() {
    $( "#menue" ).toggle();
    $( "#quickaddpage" ).hide();
    $( "#settingsinpage" ).hide();
  });

//BEGIN GO BUTTON - HIDE AND SHOW THE SEARCHBAR - AND REMOVE AFTER GOTO by SHOW FOR BETTER PLACE
$( "#go_button" ).click(function() {//OPEN
openGO()
});

//BEGIN######### BEGIN CLOSE GO BUTTON #####################

var mousedownHappened = false;

$("#g_").mousedown(function() {
  mousedownHappened = true;
});

$("#goto").blur(function() {
  if (mousedownHappened) // cancel the blur event
  {
    mousedownHappened = false;
  }
  else{
 closeGO()
  }
});


function closeGO(){
if ($(".goto").is(':visible')) {//CHEK IF OPEN
  $( "#go_button" ).toggle();
  $(".goto").css("visibility", "hidden");
  $(".goto").css("display", "none");
  $(".goto").css("opacity", "0");
  $("#clickscroll").css("display", "block");
  $("#afterGO").show();
  $(".close_button").css("display", "none");
}
}

//END########### END CLOSE GO BUTTON #####################



$(function() { //shorthand document.ready function
    $('#goto_form').on('submit', function(e) { //use on if jQuery 1.7+
        e.preventDefault();  //prevent form from submitting

        var data = $("#goto").val();
          window.location.href = "/s/" + data

        console.log(data); //use the console for debugging, F12 in Chrome, not alerts
    });
});


});//DOCUMENT READY
