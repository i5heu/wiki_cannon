$( document ).ready(function() {

  $('#header_option_gear').click(function () {
    defaultModalShow()
    $("#DefaultModalContent").html(`<button id="OptionDarkMode">Togle DarkMode</button>`)
    ToogleMenue();
  })


  $('#DefaultModal').on( 'click', '#OptionDarkMode', function () {
    alert("darktemplate = " + ToggleCockie("darktemplate"))
  })

});
