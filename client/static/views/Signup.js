export default class {
    constructor(params) {
        this.params = params;
    }

    setTitle(title) {
        document.title = title;
    }
    test(){
        console.log('hi yopta!')
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
