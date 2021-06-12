export default class {
    constructor(params) {
        this.params = params;
    }

    setTitle(title) {
        document.title = title;
    }
    init() {

        document.getElementById('signin').onclick = async function() {

            let e = document.getElementById("email").value;
            let p = document.getElementById("password").value;

            let user = {
                email: e,
                password: p,
            };
            console.log(user, "user")
            let response = await fetch('http://localhost:6969/api/signin', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(user)
            });
            let result = await response.json();
            // redirect signin page 
            if (result !== '') {
                console.log(result)

                window.location.replace('http://localhost:6969/profile')
            } else {
                //show wrong error message
                console.log(result, 'error')
            }
        }
    }
    async getHtml() {
        let wrapper = `
        <div>
        <input type='email' id='email' placeholder='email' required>
        <input type="password" id="password" placeholder='password' required>
        <input type='submit' id='signin' value="signin"/>
        </div>
        `
        return wrapper;
    }
}