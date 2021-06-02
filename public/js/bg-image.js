function ChangeIt() {
    const images = [
        'url("/img/banner1.jpg")',
        'url("/img/banner2.jpg")',
        'url("/img/banner3.jpg")',
        'url("/img/banner6.jpg")'
    ]
    const section = document.querySelector(".image")
    section.style.backgroundImage = "linear-gradient(\n" +
        "            to bottom,\n" +
        "            rgba(255, 255, 255, 0) 70%,\n" +
        "            #f6f5f7 100%\n" +
        "    )," + images[Math.floor(Math.random() * images.length)];
}
//Check with team
// setInterval(ChangeIt, 1000)


