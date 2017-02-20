$(function() {
  var refreshButtons = function($el) {
    var $rows = $el.find('.form-row');
    if ($rows.length && $rows.length == 1) {
      $el.find('.minus').hide();
      $el.find('.plus').show();
    } else if ($el.hasClass('buttons-bottom')) {
      $rows.find('.plus, .minus').hide();
      $rows.last().find('.plus, .minus').show();
    } else {
      $rows.find('.plus, .minus').show();
    }
  };

  var bindDateTimePicker = function($el) {
    $el.find('.datetimepicker').datetimepicker({ format: 'DD/MM/YYYY HH:mm', sideBySide: true });
  };

  var $formTable = $('.form-table');
  $formTable.each(function(i, el) {
    refreshButtons($(el));
    bindDateTimePicker($(el));
  });

  $formTable.on('click', '.plus', function() {
    var $currentFormTable = $(this).parents('.form-table').first();
    var $lastRow = $currentFormTable.find('.form-row').last();
    var $newRow = $lastRow.clone();
    $newRow.find('input').not('.hidden').val('');
    $lastRow.after($newRow);
    bindDateTimePicker($newRow);
    refreshButtons($currentFormTable);
    return false;
  });

  $formTable.on('click', '.minus', function() {
    var $currentFormTable = $(this).parents('.form-table').first();
    $(this).parents('.form-row').remove();
    refreshButtons($currentFormTable);
    return false;
  });
});
