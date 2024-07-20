export function SignIn({page}) {
    return (
      <div className='page' style={{display: page == "signin" ? "block" : "none"}}>
        <p>Sign in page.</p>
      </div>
    )
}
  