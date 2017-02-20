/* global: $ */

function configureDateTimeSplitInput($splitPicker) {
  $splitPicker.find('.datetimesplitpicker').datetimepicker({
    icons: {
      date: 'fa fa-calendar',
      time: 'fa fa-clock-o',
      up: 'fa fa-chevron-up',
      down: 'fa fa-chevron-down',
      previous: 'fa fa-chevron-left',
      next: 'fa fa-chevron-right',
      today: 'fa fa-crosshairs',
      clear: 'fa fa-trash-o',
      close: 'fa fa-times'
    },
    extraFormats: ['YYYY-MM-DD[T]HH:mm:ssZ']
  });

  $splitPicker.on('dp.change', function (e) {
    var $this = $(this);
    var $field = $this.find('input[type=hidden]');
    var $picker = $(e.target);
    var current = moment.parseZone($field.val());
    var picked = e.date;
    if ($picker.hasClass('date')) {
      current.year(picked.year()).month(picked.month()).date(picked.date());
    } else if ($picker.hasClass('time')) {
      current.hours(picked.hours()).minutes(picked.minutes()).seconds(picked.seconds());
    }
    $field.val(current.format());
  });

  $splitPicker.closest('form').submit(function () {
    // disable custom fields so they're not submitted to the server
    $splitPicker.find('.datetimesplitpicker > input').attr('disabled', true);
  });
}

$(function () {
  $('.date-time-split-picker').each(function () {
    $input = $(this);
    configureDateTimeSplitInput($input);
  });
});
