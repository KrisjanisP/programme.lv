import 'bootstrap/dist/css/bootstrap.css'
import '../styles/globals.css'
import Head from "next/head";

import {useEffect} from 'react'

function MyApp({Component, pageProps}) {
    useEffect(() => {
        require("bootstrap/dist/js/bootstrap");
    }, [])

    return (
        <>
            <Head>
                <meta name="viewport" content="width=device-width, initial-scale=1"/>
                <title>programme.lv</title>
            </Head>
            <Component {...pageProps} />
        </>
    );
}

export default MyApp
