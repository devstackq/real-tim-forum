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

        let user = {
            email: "",
            password: "",
        };
        user = super.fillObject(user)
        
        let result = await super.fetch('signin', user)
        if (result !== null) {
            localStorage.setItem('isAuth', true)
            window.location.replace('http://localhost:6969/profile')
        } else {
            localStorage.setItem('isAuth', false)
            super.showNotify('incorrect login or password', 'error')
        }
    }

    init() {
        localStorage.setItem('isAuth', false)
        document.querySelector('#signin').onclick = this.signin
    }

    async getHtml() {
        let body = `
        <div>
        <input type='email' id='email' placeholder='email' required>
        <input type="password" id="password" placeholder='password' required>
        <input type='submit' id='signin' value="signin"/>
        </div>
        `
        let header = super.showHeader('free');
        return header + body
    }
};