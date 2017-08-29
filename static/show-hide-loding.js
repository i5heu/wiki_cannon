
$( document ).ready(function() {

  LoadingPoints();

  $("#DefaultModalClose").mousedown(function() {
    pwdhash = 0
    defaultModalHide();
  });

});

function defaultModalHide(){
  $("#DefaultModalContainer").hide()
  $("#DefaultModalContent").html(`MODAL SHOLD BE CLOSED NOW`)
}

function defaultModalShow(){
  $("#DefaultModalContainer").show()
  $("#DefaultModalContent").html(`LODING <span class="LodingIndicator"></span>`)
}


function LoadingPoints(){ // Makes the Loading Points ... .. .
  var LodingIndicatorCunter = 0;
  var LodingIndicatorPoint = "&emsp;";
  $(".LodingIndicator").html("&emsp;");

  window.setInterval(function(){
    $(".LodingIndicator").html(LodingIndicatorPoint);
    LodingIndicatorPoint = LodingIndicatorPoint.concat(".")


    LodingIndicatorCunter++
    if (LodingIndicatorCunter > 3) {
      LodingIndicatorPoint = "&emsp;"
      LodingIndicatorCunter = 0
    }

    if (LodingIndicatorCunter == 1) {
      LodingIndicatorPoint = "."
    }

  }, 600);
}
