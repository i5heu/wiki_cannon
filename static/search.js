//SEARCH WIKI_CANNON

function GetSearchTerm(){
  return $("#searchterm").text();
}

function SearchTextSearch(){
  data = '{"PWD":"'+  $.cookie("pwd") + `", "SearchValue":"` + GetSearchTerm() + '"}';

  $.ajax({
              type:"POST",
              url: "/api-search/",
              data:data,
              success: function (response){
                    $(".flexparent").append('<div class="SearchFlexChild"></div>')
                    var json = $.parseJSON(response);
                    console.log(GetSearchTerm());
                    $(json.SearchResult).each(function(index, item) {
                        $('.SearchFlexChild').append(item.SR-ID);
                    });

                  }
        });


}

$( document ).ready(function() {
  if(window.location.href.indexOf("/s/") > -1) {
    SearchTextSearch();
    }
}); // DO NOT REMOVE DOC RDY
