import { useEffect, useState } from 'react'
import env from './env.js'

const sleep = ms => new Promise(r => setTimeout(r, ms));

async function updateStatus(message) {
  let statusBox = document.getElementById("status-box");
  statusBox.innerText = message
  statusBox.style.opacity = 1.5;
  for (let i = 0; i < 15; i++) {
    await sleep(100)
    statusBox.style.opacity -= 0.1
  }
  statusBox.innerText = " "
}


export function SignUp({page}) {
  console.log(env)
  const [email, setEmail] = useState("")
  const [name, setName] = useState("")

    function submitForm(e) {
        e.preventDefault();
        console.log(JSON.stringify({
          email: email,
          name: name,
        }))

        fetch("http://" + env.baseUrl + "/users", {
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
          (response) => {
            if (response.status !== 201) {
              if (response.status === 409) {
                updateStatus("Email already used in site.")
                console.log("Email already signed up.")
              } else {
                updateStatus("Unknown error: " + response.statusText)
                console.log("Unknown response code of " + response.status + ". Response text: " + response.statusText)
              }
            }
            return response.json()
          }
        ).then(
          data => console.log(data)
        ).catch(
          error => console.error(error)
        );


    }
    return (
      <div className='page' style={{display: page === "signup" ? "block" : "none"}}>
        <form onSubmit={submitForm}>
          <table><tbody>
            <tr>
              <th>Name</th>
              <td>
              <input type='text' name='email' onChange={(e) => setName(e.target.value)}/>
              </td>
            </tr>
            
            <tr>
              <th>Email</th>
              <td>
                <input type='text' name='name' onChange={(e) => setEmail(e.target.value)}/>
              </td>
            </tr>

            <tr>
              <td></td>
              <td>
              <button type='submit' name='submit' value='Sign Up'>
                Sign Up
            </button>
              </td>
            </tr>
            <tr><td colSpan={2} className='message-row'><span id='status-box'> </span></td></tr>
            </tbody></table>
        </form>
      </div>
    )
}
