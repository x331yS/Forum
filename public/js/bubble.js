const contain = document.querySelector('.fun')
for (let i = 1; i <= 300; i++) {
    const blocks = document.createElement('div')
    blocks.classList.add('block')
    blocks.onclick=generate
    contain.appendChild(blocks)
}
function generate() {
    anime({
        targets: '.block',
        translateX: function () {
            return anime.random(-1100, 1100)
        },
        translateY: function () {
            return anime.random(-100, 8000)
        },
        scale: function () {
            return anime.random(1, 5)
        }
    })
}
generate()