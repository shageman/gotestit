# Go Test It [![Build Status](https://travis-ci.org/shageman/gotestit.png?branch=master)](https://travis-ci.org/shageman/gotestit)

Comparison of go lang testing libraries

## Libraries in comparison

*   testing: http://golang.org/pkg/testing/  
	Last Activity: -  
	
*   testify: https://github.com/stretchr/testify/  
	Last Activity: 2 days ago on 7/3/2013
	
*   gocheck: http://labix.org/gocheck  
	Last Activity: 16 days ago on 7/3/2013
	
*   prettytest: https://github.com/remogatto/prettytest  
	Last Activity: 60 days ago on 7/3/2013
	
*   go-spec: https://github.com/bmatsuo/go-spec  
	Last Activity: 700 days ago on 7/3/2013
	
*   gospec: https://github.com/orfjackal/gospec  
	Last Activity: 350 days ago on 7/3/2013

*   mao: https://github.com/azer/mao  
    Last Activity: 9 days ago on 7/9/2013

*   zen: https://github.com/pranavraja/zen  
    License: Apache  
    Last Activity: 2 days ago on 7/13/2013

## Assertions/Matchers

<table>
    <tbody>
    <tr>
        <th>Name</th>
        <th>Testing</th>
        <th>testify</th>
        <th>gocheck</th>
        <th>prettytest</th>
        <th>go-spec</th>
        <th>gospec</th>
        <th>mao/zen</th>
    </tr>
    <tr>
        <td>License</td>
        <td>BSD</td>
        <td>MIT</td>
        <td>BSD</td>
        <td>MIT</td>
        <td>BSD</td>
        <td>Apache</td>
        <td>MIT/Apache</td>
    </tr>
    <tr>
        <td>Assertions</td>
        <td></td>
        <td></td>
        <td></td>
        <td>uses gocheck</td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>Style</td>
        <td>make your own</td>
        <td>assert</td>
        <td>spec</td>
        <td>spec</td>
        <td>spec</td>
        <td>spec</td>
        <td>spec</td>
    </tr>
    <tr>
        <td>Equal</td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
    </tr>
    <tr>
        <td>IsSame</td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td>✓</td>
        <td></td>
    </tr>
    <tr>
        <td>DeepEqual</td>
        <td></td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>True</td>
        <td></td>
        <td>✓</td>
        <td></td>
        <td></td>
        <td></td>
        <td>✓</td>
        <td></td>
    </tr>
    <tr>
        <td>False</td>
        <td></td>
        <td>✓</td>
        <td></td>
        <td></td>
        <td></td>
        <td>✓</td>
        <td></td>
    </tr>
    <tr>
        <td>Nil</td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
    </tr>
    <tr>
        <td>Empty</td>
        <td></td>
        <td>✓</td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>Error</td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>Implements</td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>IsType</td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>StringContains</td>
        <td></td>
        <td>✓</td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>StringMatches</td>
        <td></td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>Collection</td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td>✓</td>
        <td></td>
    </tr>
    <tr>
        <td>Panics</td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>HasLen</td>
        <td></td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>Matches</td>
        <td></td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>Satisfy</td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
        <td></td>
    </tr>
    <tr>
        <td>Within</td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td>✓</td>
        <td></td>
    </tr>
    </tbody>
</table>

*CollectionContains allows checks for All, Any, Exactly, InOrder, and InPartialOrder

## Sources

https://code.google.com/p/go-wiki/wiki/Projects#Testing
