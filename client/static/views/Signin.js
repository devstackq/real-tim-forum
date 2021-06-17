import Parent from "./Parent.js";

export default class Signin extends Parent {

    constructor(text, type, params) {
        super(text, type)
        this.params = params;
    }

    setTitle(title) {
        document.title = title;
    }

    async signin() {

        let e = document.getElementById("email").value;
        let p = document.getElementById("password").value;

        let user = {
            email: e,
            password: p,
        };

        let response = await fetch('http://localhost:6969/api/signin', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json;charset=utf-8'
            },
            body: JSON.stringify(user)
        });
        // redirect signin page 
        if (response.status == 200) {
            super.showNotify('', 'hide')
            localStorage.setItem('isAuth', true)
            window.location.replace('http://localhost:6969/profile')
        } else {
            localStorage.setItem('isAuth', false)
            super.showNotify('incorrect login or password', 'error')
        }
    }

    init() {

        let btn = document.querySelector('#signin')
        btn.onclick = this.signin
            //send data to Parent constructoer -> then use parent method value
            // let n = new Signin('soma', 'error')
    }

    async getHtml() {
        let wrapper = `
        <div>
        <input type='email' id='email' placeholder='email' required>
        <input type="password" id="password" placeholder='password' required>
        <input type='submit' id='signin' value="signin"/>
        </div>
        `
        let h = super.showHeader('free');
        // console.log(h, 'hed')
        return h + wrapper
    }
};