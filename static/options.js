$( document ).ready(function() {

  $('#header_option_gear').click(function () {
    defaultModalShow()
    $("#DefaultModalContent").html(`<button id="OptionDarkMode">Togle DarkMode</button>`)
    $( "#menue" ).hide();
  })


  $('#DefaultModal').on( 'click', '#OptionDarkMode', function () {
    alert("darktemplate = " + ToggleCockie("darktemplate"))
  })

});
