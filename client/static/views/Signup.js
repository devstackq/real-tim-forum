export default class {
    constructor(params) {
        this.params = params;
    }

    setTitle(title) {
        document.title = title;
    }
    init() {

        document.getElementById('signup').onclick = async function() {

            let e = document.getElementById("email").value;
            let p = document.getElementById("password").value;
            let u = document.getElementById("username").value;
            let f = document.getElementById("fName").value;
            let a = document.getElementById("age").value;
            let c = document.getElementById("city").value;
    
            let user = {
                email: e,
                password: p,
                username: u,
                fullname: f,
                age: a,
                city: c
            };
            console.log(user, "user")
            let response = await fetch('http://localhost:6969/signup', {
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
                window.location.replace('http://localhost:6969/signin')
            }
        }
     }
    async getHtml() {
        let wrapper = `
        <div>
        <input type="text" id='fName' placeholder='full name'>
        <input type='email' id='email' placeholder='email'>
        <input type="text" id='username' placeholder='nick'>
        <input type="password" id="password" placeholder='password'>
        <input type="number" id='age' placeholder='age'>
        <input type="text" id='city' placeholder='city'>
        <input type='submit' id='signup' value="register"/>
        </div>
        `
        return wrapper;
    }
}