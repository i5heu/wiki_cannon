function defaultModalHide(){
  $("#DefaultModalContainer").hide()
  $("#DefaultModalContent").html(`MODAL SHOLD BE CLOSED NOW`)
}

function defaultModalShow(){
  $("#DefaultModalContainer").show()
  $("#DefaultModalContent").html(`LODING <span class="LodingIndicator"></span>`)
}


$( document ).ready(function() {
  var LodingIndicatorCunter = 0;
  var LodingIndicatorPoint = "";
  $(".LodingIndicator").html("");

  window.setInterval(function(){
    $(".LodingIndicator").html(LodingIndicatorPoint);
    LodingIndicatorPoint = LodingIndicatorPoint.concat(".")

    LodingIndicatorCunter++
    if (LodingIndicatorCunter > 3) {
      LodingIndicatorPoint = ""
      LodingIndicatorCunter = 0
    }
  }, 600);
});
