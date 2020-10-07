import React from 'react';
import {ApolloProvider} from "@apollo/client";

export default function Layout() {
        return (
            <div className='layout'>
                <div className="overlay"></div>
                <div className="scanline"></div>
            </div>
        );
}