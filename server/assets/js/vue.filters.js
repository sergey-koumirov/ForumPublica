var frm = new Intl.NumberFormat('ru-RU', {
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
});

var frmFl = new Intl.NumberFormat('ru-RU', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
});


Vue.filter('number', function (value) {
    if(value==undefined){
      return '...';
    }
    return frm.format(value);
});

Vue.filter('float', function (value) {
    if(value==undefined){
        return '...';
    }
    return frmFl.format(value);
});

Vue.filter('price', function (value) {
    if(!value && value!=0){
        return '...';
    }
    return frmFl.format(value);
});

Vue.filter('seconds', function(seconds){
    var s = seconds||0,
        ss = s % 60,
        mm = Math.trunc(s/60) % 60,
        hh = Math.trunc(s/(60*60)) % 24,
        dd = Math.trunc(s/(60*60*24));

    var result = ("0" + hh).slice(-2) + ":" + ("0" + mm).slice(-2) + ":" +("0" + ss).slice(-2);
    if( dd > 0){ result = dd + "d " + result; }
    return result;
});
