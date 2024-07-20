import { useEffect, useState } from 'react'



export function SignUp({page}) {
  const [email, setEmail] = useState("")
  const [name, setName] = useState("")

    function submitForm(e) {
        e.preventDefault();
        console.log(email)

        fetch('http://localhost:8080/users', {
          method: 'POST',
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: email,
            name: name,
          })
        }).then(
          response => response.json()
        ).then(
          data => console.log(data)
        ).catch(
          error => console.error(error)
        );
    }
    return (
      <div className='page' style={{display: page == "signup" ? "block" : "none"}}>
        <form onSubmit={submitForm}>
            <input type='text' name='email' onChange={(e) => setEmail(e.target.value)}/>
            <input type='text' name='name' onChange={(e) => setName(e.target.value)}/>
            <button type='submit' name='submit' value='Sign Up'>
                Sign Up
            </button>
        </form>
      </div>
    )
}
