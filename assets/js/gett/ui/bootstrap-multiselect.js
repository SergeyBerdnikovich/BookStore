$(function() {
  "use strict";

  $('.filters .multi select').each(function() {
    var select = $(this),
     container = select.parents('.multi'),
       options = {
         maxHeight: '270',
         includeSelectAllOption: true,
         numberDisplayed: 1,
         multiple: true,
         inheritClass: true,
         disableIfEmpty: true,
         buttonText: function(options, element) { return select.data('title'); }
       };

    if (select.hasClass('filter-enabled')) {
      $.extend(options, {
        enableFiltering: true,
        enableCaseInsensitiveFiltering: true,
        includeFilterClearBtn: true
      });
    }

    select.multiselect(options);

    var select_all = container.find('.multiselect-all input');
    if (select_all.length && select.data('all-checked')) {
      select.multiselect('selectAll', false);
    }

  });
});
