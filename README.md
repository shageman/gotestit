# Go Test It [![Build Status](https://travis-ci.org/shageman/gotestit.png?branch=master)](https://travis-ci.org/shageman/gotestit)

Summary from the July presentation: There are a couple of interesting testing libraries in go. "testing" comes with the standard library, and gives only very basic support for asserting facts. 

Of the other libraries, I recommend looking at "testify/assert", "zen", and "gocheck". testify/assert is used in conjunction with testing to which it adds a lot of useful assertions, but doesn't change the basics of how a test is written. Zen implements the best version of BDD style testing I have seen in go. Essentially, you write one test into which you embed describes, its, and expects. Like testify/assert "gocheck" comes with many matchers out of the box. Gocheck requires creating a suite struct and uses a non-standard way of running the test methods. Its most useful feature to me are setup and teardown methods to allow tests in a suite to share common code. The shared state can lead to problems with concurrent code.

As far as the other libraries go, check them out to see the different styles. Beware of go-spec: The version tested here, never fails assertions!

So, summary of the summary: if you are just trying to test stuff and need assertions, go with testify/assert. If you want to see and play with BDD for go, use zen.


[Slides from my presentation at Denver Gopher's 7/25/2013](https://github.com/shageman/gotestit/blob/master/20130725denverGophersPresentation.pdf)

Comparison of go lang testing libraries

## Libraries in comparison

*   testing: http://golang.org/pkg/testing/  
    Last Activity: -  

*   GoConvey: https://github.com/smartystreets/goconvey/  
    Last Activity: 2 days ago 10/25/2013

*   testify: https://github.com/stretchr/testify/  
    Last Activity: 7 days ago on 7/24/2013
	
*   gocheck: http://labix.org/gocheck  
    Last Activity: 30 days ago on 7/24/2013
	
*   prettytest: https://github.com/remogatto/prettytest  
    Last Activity: 90 days ago on 7/24/2013
	
*   go-spec: https://github.com/bmatsuo/go-spec  
    Last Activity: 700 days ago on 7/24/2013
	
*   gospec: https://github.com/orfjackal/gospec  
    Last Activity: 350 days ago on 7/24/2013

*   mao: https://github.com/azer/mao  
    Last Activity: 17 days ago on 7/9/2013

*   zen: https://github.com/pranavraja/zen  
    Last Activity: 14 days ago on 7/13/2013

## Assertions/Matchers

<table>
    <tbody>
    <tr>
        <th>Name</th>
        <th>Testing</th>
        <th>GoConvey</th>
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
        <td></td>
        <td>uses gocheck</td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>Style</td>
        <td>make your own</td>
        <td>spec</td>
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
        <td>✓</td>
    </tr>
    <tr>
        <td>IsSame</td>
        <td></td>
        <td>✓</td>
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
        <td>✓</td>
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
        <td>✓</td>
        <td></td>
        <td>✓</td>
        <td>✓</td>
    </tr>
    <tr>
        <td>Empty</td>
        <td></td>
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
        <td></td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td>✓</td>
        <td></td>
    </tr>
    <tr>
        <td>Implements</td>
        <td></td>
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
        <td>✓</td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>StringContains</td>
        <td></td>
        <td>✓</td>
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
        <td>✓</td>
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
        <td>✓</td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td>HasLen</td>
        <td></td>
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
        <td></td>
        <td>✓</td>
        <td></td>
    </tr>
    </tbody>
</table>

*CollectionContains allows checks for All, Any, Exactly, InOrder, and InPartialOrder

## Sample tests

The code to be tested is taken from the [cloudfoundry loggregator project](https://github.com/cloudfoundry/loggregator). Specifically, we are testing a piece that receives data on a channel and sends it off via UPD. We will be testing that [the code](https://github.com/shageman/gotestit/blob/master/src/gotestit/loggregatorclient/loggregatorclient.go) sends data, but that it does not send any empty data it might receive.

All the tests are in packages relating to their name within the [gotestit package](https://github.com/shageman/gotestit/tree/master/src/gotestit).

## Sources

https://code.google.com/p/go-wiki/wiki/Projects#Testing
