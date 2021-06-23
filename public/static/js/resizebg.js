var script1= document.createElement('script');
var script2= document.createElement('script');
script1.src = '../static/js/bg-image.js';
script2.src = '../static/js/phone-bg.js';
script1.type = 'text/javascript';
script2.type = 'text/javascript';
if (!(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent))) {
    document.getElementsByTagName('body')[0].appendChild(script1);
} else {
    document.getElementsByTagName('body')[0].appendChild(script2);
}
