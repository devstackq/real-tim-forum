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
            let response = await fetch('http://localhost:6969/signin', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(user)
            });
            let result = await response.json();
            // redirect signin page 
            console.log(result, 'result')
            if (result > 0 && result != undefined) {
                // window.location.replace('http://localhost:6969/signin')
            }
        }
    }
    async getHtml() {
        let wrapper = `
        <div>
        <input type='email' id='email' placeholder='email'>
        <input type="password" id="password" placeholder='password'>
        <input type='submit' id='signin' value="signin"/>
        </div>
        `
        return wrapper;
    }
}