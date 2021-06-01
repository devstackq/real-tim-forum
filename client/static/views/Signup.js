const Signup = async() => {

    let e = document.getElementById("email").value;
    let p = document.getElementById("password").value;

    let user = {
        email: e,
        password: p
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
}