console.log('Start APP');
let plant = new Plant();
const plantHTML = document.getElementById("plantName");

// Titel extrahieren
function readTitleName(text, substring) {
    const lines = text.split('\n')
    const regexp = new RegExp(`${substring}"(?<f>.*)" `)
    let res = "none"
    if (lines.length > 0) {
        lines.forEach(l => {
            if (l.indexOf(substring) > -1) {               
                const matches = l.match(regexp)
                if (matches) {
                    res = matches.groups.f
                }
            } 
        })
    }
    return res
}

// Menulinks mit der Anlage ID versehen
function createNavigation(id) {
    const leftNavs = document.querySelectorAll(".header-menu__left li a");
    leftNavs.forEach((n) => {
    let attr = n.getAttribute("href");
    n.setAttribute("href", attr + "/" + id);
    });
}