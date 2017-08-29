function WikiLoadingON(){
  var LodingIndicatorCunter = 0;
  var LodingIndicatorPoint = "&emsp;";
  $("#LoadingIndicator").html("&emsp;");

  window.setInterval(function(){
    $("#LoadingIndicator").html(LodingIndicatorPoint);
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

$( document ).ready(function() {
  WikiLoadingON();
}); // DO NOT REMOVE DOC RDY
