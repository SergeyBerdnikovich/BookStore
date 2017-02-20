$.fn.dataTableExt.sErrMode = function(settings, tn, msg) {
  console.log("Error loading data: " + msg);
};

$(function() {
  var Datatableable = (function(){

    function init(datatables) {
      $.each(datatables, function(i, t){
        var table = $(t);
        // check Datatable already initialized
        if (!$.fn.dataTable.isDataTable(table)) {
          table.DataTable($.extend({
            dom: "<'row'<'col-xs-6'l><'col-xs-6'f>r>t<'row'<'col-xs-1'i><'col-xs-10'p>>",
            pagingType: "full_numbers",
            info: false,
            lengthChange: false
          }, ($(table).data('datatable') || {})));
        }

        $('body').trigger('remoteLink', table);
      });
    }

    return init;
  })();

  var datatables = $("table[data-datatable]");
  if (datatables.length > 0) {
    Datatableable(datatables);
  }
});
