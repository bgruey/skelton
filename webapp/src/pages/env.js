const env = {
    "host": process.env.REACT_APP_BACKEND_HOST,
    "port": process.env.REACT_APP_BACKEND_PORT
}

env.baseUrl = env.host + ":" + env.port

export default env