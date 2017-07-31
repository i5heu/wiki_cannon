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
                    $(".flexparent").append('<div class="SearchFlexChild" id="TextSearch"></div>')
                    var json = $.parseJSON(response);
                    console.log(GetSearchTerm());
                    $(json.SearchResult).each(function(index, item) {
                        $('#TextSearch').append("<div>#" + item.id +`  <a href="` + item.namespace + "/" +item.title + `"> `+item.namespace + "/" +item.title+`</a>  [` +  item.tags  + `] >> <p class="lightgreybackgorund">`+ item.text +" </p></div>");
                    });

                  }
        });


}

$( document ).ready(function() {
  if(window.location.href.indexOf("/s/") > -1) {
    SearchTextSearch();
    }
}); // DO NOT REMOVE DOC RDY
