class Alert {

    static show ({message, type='success'}) {
        const a = new Alert()
        const alertBox = document.getElementById('login-alert')
        alertBox.classList.add(type, 'alert')
        alertBox.innerHTML = a.success + `           
            <p>${message}</p>
        `
    }

    success =/*css*/`<style>
                .alert {
                    border: solid 5px green;
                    background-color: lightgreen;
                    width: 100vw;
                }
                .alert>p {
                    color: green;
                }
            </style>
            `
    
}