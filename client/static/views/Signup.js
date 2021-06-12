export default class {
    constructor(params) {
        this.params = params;
    }
    setTitle(title) {
        document.title = title;
    }
    init() {
        //post query
        document.getElementById('signup').onclick = async function() {

            let e = document.getElementById("email").value;
            let p = document.getElementById("password").value;
            let u = document.getElementById("username").value;
            let f = document.getElementById("fName").value;
            let a = document.getElementById("age").value;
            let c = document.getElementById("city").value;
            let g = document.getElementById("gender").value;

            let user = {
                email: e,
                password: p,
                username: u,
                fullname: f,
                age: a,
                city: c,
                gender: g,
            };

            console.log(user, "user")
            let response = await fetch('http://localhost:6969/api/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(user)
            });
            let result = await response.json();
            // redirect signin page 
            if (result > 0 && result != undefined) {
                window.location.replace('http://localhost:6969/signin')
            } else {
                //show notiy
                console.log(result)
            }
        }
    }
    async getHtml() {
        let wrapper = `
        <div>
        <input type="text" id='fName' required="true" placeholder='full name'>
        <input type='email' id='email' required placeholder='email'>
        <input type="text" id='username' required placeholder='nick'>
        <input type="password" id="password" required placeholder='password'>
        <input type="number" id='age' required placeholder='age'>
      <label> gender
        <select id='gender' placeholder='gender'>
        <option></option>
        <option>man</option>
        <option>woman</option>
      </select>
      </label>
        <input type="text" id='city' required placeholder='city'>
        <input type='submit' id='signup' value="register"/>
        </div>
        `
        return wrapper;
    }
}