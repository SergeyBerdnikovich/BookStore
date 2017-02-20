$(function() {
  var ResourceSelect = (function(){
    function init(element) {

      var defaults = {per_page: 20, url: null, current: {}}
      var select   = $(element);
      var options  = $.extend(defaults, select.data('resources'));

      select.select2({
        minimumInputLength: 3,
        dropdownAutoWidth: true,
        allowClear: true,

        // NOTE: https://select2.github.io/announcements-4.0.html#removed-initselection
        initSelection: function (element, callback) {
          if (options.current.id && options.current.text) {
            callback(options.current);
          }
        },

        ajax: {
          url: options.url,
          dataType: 'json',
          cache: true,

          data: function (term, page) {
            return {
              query: term, // try server-side filtering
              page: page,
              per_page: options.per_page,
              format: 'json',
              current: select.val()
            };
          },
          results: function (data, page, query) {
            var results = [];
            var more = (page * options.per_page) < data.total;

            $.each(data.data, function(i, resource) {
              if (resource[options.by].toUpperCase().indexOf(query.term.toUpperCase()) >= 0) {
                results.push({id: resource.id, text: resource[options.by]});
              }
            });

            return { results: results, more: more };
          }
        }
      });
    }

    return init;
  })();

  var selects = $("*[data-resources]");

  if (selects.length > 0) {
    $.each(selects, function(i,s) {
      ResourceSelect(s)
    });
  }

});
