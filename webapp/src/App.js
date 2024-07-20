import { useEffect, useState } from 'react';

import { About } from './pages/about.js';
import { SignIn } from './pages/signin.js'
import { SignUp } from './pages/signup.js'

import './App.css';


function Header({page, setPageClick}) {
  return (
    <>
    <h2 className='header'>TAF: Tulsa Awards Fellowship</h2>
    <span  onClick={() => setPageClick("about")} className={page === "about" ? "menu-clicked" : "menu"}>
        <b>About</b>
      </span>
      <span onClick={() => setPageClick("signin")} className={page === "signin" ? "menu-clicked" : "menu"}>
        <b>Sign In</b>
      </span>
      <span onClick={() => setPageClick("signup")} className={page === "signup" ? "menu-clicked" : "menu"}>
        <b>Sign Up</b>
      </span>
    </>
    
  )
}


function Footer() {
  return (
    <div>
      <p>&copy; Never by No One</p>
    </div>
  )
}

function App() {
  const [squares, setSquares] = useState(Array(9).fill(null))

  function handleClick(i) {
    const nextSquares = squares.slice();
    nextSquares[i] = squares[i] ? null : "X"
    setSquares(nextSquares);
  }

  const [page, setPage] = useState("about")
  function menuClick(name) {
    setPage(name)
  }

  return (
    <div className='main'>
      <Header page={page} setPageClick={menuClick} />

      <About page={page} />

      <SignIn page={page} />

      <SignUp page={page} />

      <Footer />
    </div>
  )
}

export default App;
