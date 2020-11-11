package main

import (
	"encoding/xml"
	"fmt"
)

type BookXML struct {
	XMLName   	xml.Name    		`xml:"book"`                    // used to make sure this struct is marshalled as <book> node, don't need to set value in go
	Version   	string      		`xml:"version,attr,omitempty"`  // xml format version for this xml
	Tutorial  	bool       		`xml:"tutorial,attr,omitempty"` // if xml is tutorial
	Bundles   	[]BundleXML 		`xml:"bundles>bundle"`          // bundle list
	Quizs 		[]QuizXML 		`xml:"quizMaterials"`
	Remaining 	[]XmlNode  		`xml:",any"`                    // other nodes in the xml (don't need to parse in go)
}

// BookXML contains corresponding content for bundle node in book xml
type BundleXML struct {
	ID      string `xml:"id,attr"`
	Type    string `xml:"type,attr"`
	URL     string `xml:"url,attr"`
	Version string `xml:"version,attr"`
}

// XmlNode is used to hold arbitrary xml node content
type XmlNode struct {
	XMLName    xml.Name  			// save node name from unmarshal, so ti can be marshalled to the same node name
	Text       string     `xml:",innerxml"` // inner xml of the node
	Attributes []xml.Attr `xml:",any,attr"` // all attributes of the node (preserving order)
}

type QuizXML struct {
	XMLName    xml.Name   `xml:"quiz"` 	// save node name from unmarshal, so ti can be marshalled to the same node name
	Text       string     `xml:",innerxml"` // inner xml of the node
	Attributes []xml.Attr `xml:",any,attr"` // all attributes of the node (preserving order)
}

func main() {
	var book BookXML
	
	bytes := []byte(`<book version="2">
				<textures>
    					<quizTexture id="76503ce2-9fc7-4abf-a7ec-eeef120750fb" isText="true">
      						<metaInfo><![CDATA[問題1]]></metaInfo>
      						<metaInfo><![CDATA[This is question 1]]></metaInfo>
      						<optionText><![CDATA[This is option 1]]></optionText>
      						<optionText><![CDATA[This is option 2]]></optionText>
      						<optionText><![CDATA[This is option 3]]></optionText>
      						<optionText><![CDATA[This is option 4]]></optionText>
    					</quizTexture>
    					<quizTexture id="d7e42a94-1a45-4d5b-9f6a-502d8c5f8d64" isText="false">
      						<metaInfo><![CDATA[問題2]]></metaInfo>
      						<metaInfo><![CDATA[This is question 2]]></metaInfo>
      						<metaInfo><![CDATA[詳解]]></metaInfo>
      						<metaInfo><![CDATA[This is correct answer 2]]></metaInfo>
      						<optionImage url="e0096e03-52dd-4280-aec3-98d65f19bbc6" />
      						<optionImage url="c275d3bd-70a8-4b8c-b75d-22424ee342d3" />
    					</quizTexture>
  				</textures>
				<quizMaterials>
					<quiz id="99be158a-2632-4eed-910c-47ca2519ea21" name="" isText="True" isMultiSelect="False" selectedAction="JumpAnyway">
      						<question><![CDATA[This is question 1]]></question>
      						<detailed><![CDATA[]]></detailed>
      						<actions>
        						<actionType action="JumpAnyway" type="next_page" page="" />
        						<actionType action="JumpWhenCorrect" type="next_page" page="" />
        						<actionType action="JumpWhenWrong" type="next_page" page="" />
      						</actions>
      						<options>
        						<option id="2934627c-a2f1-4cf9-9779-ac6f9b2b0e97" image="" isCorrect="True">
          							<description><![CDATA[This is option 1]]></description>
        						</option>
        						<option id="50e04e21-1ae1-4e85-9d82-8e908e2a20e2" image="" isCorrect="False">
         							<description><![CDATA[This is option 2]]></description>
        						</option>
        						<option id="8790e8a9-6953-47dd-9f82-b8c37f0359d7" image="" isCorrect="False">
          							<description><![CDATA[This is option 3]]></description>
        						</option>
        						<option id="89275f81-8ea0-4509-b41e-377c2b42efdb" image="" isCorrect="False">
          							<description><![CDATA[This is option 4]]></description>
        						</option>
      						</options>
    					</quiz>
					<quiz id="0afac3fa-72f4-4921-9df1-9f18b229d5ab" name="" isText="False" isMultiSelect="True" selectedAction="JumpWhenWrong">
      						<question><![CDATA[This is question 2]]></question>
      						<detailed><![CDATA[This is correct answer 2]]></detailed>
      						<actions>
        						<actionType action="JumpAnyway" type="next_page" page="" />
        						<actionType action="JumpWhenCorrect" type="next_page" page="" />
        						<actionType action="JumpWhenWrong" type="prev_page" page="" />
      						</actions>
      						<options>
        						<option id="be8e3a27-141a-44cc-b27e-12e5ccdab736" image="e0096e03-52dd-4280-aec3-98d65f19bbc6" isCorrect="True">
          							<description><![CDATA[]]></description>
        						</option>
        						<option id="90f641d2-b95e-4dba-9bbf-a1fd23c629fc" image="c275d3bd-70a8-4b8c-b75d-22424ee342d3" isCorrect="False">
          							<description><![CDATA[]]></description>
        						</option>
      						</options>
    					</quiz>
				</quizMaterials>
			</book>`)
	
	if err := xml.Unmarshal(bytes, &book); err != nil {
		fmt.Printf("Error: %v", err)
	}
	
	

	fmt.Printf("%+v", book)
}
