class Alert {

    static show ({message, type='success'}) {
        const a = new Alert()
        const alertBox = document.getElementById('login-alert')
        alertBox.classList.add(type, 'alert')
        if (type === 'error') {
            alertBox.innerHTML = a.error + `
                <p>${message}</p>
            `
        } else if(type === 'success') {
        alertBox.innerHTML = a.success + `           
            <p>${message}</p>
        `
        } else if(type === 'warning') {
            alertBox.innerHTML = a.warning + `
                <p>${message}</p>
            `
        }
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
    error =/*css*/`<style>
        .alert {
            border: solid 5px red;
            background-color: lightcoral;
            width: 100vw;
        }
        .alert>p {
            color: red;
        }
        </style>
        `
    warning =/*css*/`<style>
        .alert {
            border: solid 5px orange;
            background-color: beige;
            width: 100vw;
        }
        .alert>p {
            color: orange;
        }
        </style>
        `
    
}