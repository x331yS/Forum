function ChangeIt() {
    const images = [
        'url("/img/banner1phone.jpg")',
        'url("/img/banner2phone.jpg")',
        'url("/img/banner3phone.jpg")',
        'url("/img/banner4phone.jpg")',
        'url("/img/banner5phone.jpg")'
        // 'url("/img/banner6phone.jpg")',
        // 'url("/img/banner7phone.jpg")',
        // 'url("/img/banner8phone.jpg")'
    ]
    const section = document.querySelector(".image")
    section.style.backgroundImage = "linear-gradient(\n" +
        "            to bottom,\n" +
        "            rgba(255, 255, 255, 0) 70%,\n" +
        "            #333333 100%\n" +
        "    )," + images[Math.floor(Math.random() * images.length)];
}