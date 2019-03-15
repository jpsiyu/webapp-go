import React from 'react'

class Welcome extends React.Component{
    constructor(props){
        super(props)
    }

    render(){
        return <div className='welcome'>
            <h1>Welcome to use this framework!</h1>
            <img src="/images/go.png" />
        </div>
    }
}

export default Welcome