from key import Bard_Key as BardKey
from key import OpenAi_Key as OpenAiKey
import google.generativeai as palm
import openai
import re
import os

def generate_tests(LLM, filename, language):
    code = """"""
    if LLM == "Bard":
        tests = bard_tests(BardKey, code, language)
        return tests
    elif LLM == "OpenAi":
        tests = openai_tests(OpenAiKey, code, language)
    else:
        print("Invalid LLM")
        return None

def bard_tests(GptKey ,code, language):
    model = "models/text-bison-001"
    palm.configure(api_key=GptKey)

    prompt = f"Generate test cases file for the code: {code}\n"
    postPromt = f"{{type: smart contract; Language: {language};}}"

    prompt = prompt+postPromt

    completion = palm.generate_text(
        model=model,
        prompt=prompt,
        temperature=0,
        max_output_tokens=2048,
    )
    Message = completion.result
    
    return Message



def openai_tests(GptKey ,code, language):
    Model = 'gpt-3.5-turbo'
    openai.api_key = GptKey

    prompt = f"Generate test cases file for the code: {code}\n"
    postPromt = f"{{type: smart contract; Language: {language};}}"
    
    prompt = prompt+postPromt

    response=openai.ChatCompletion.create(
        model=Model,
        messages=[{"role":"user",
                "content":prompt}
                ],
        max_tokens=2048,                                                                                                                                                           
    )

    Message = response.choices[0].message.content

    return Message


def Create_testfile(language, filename, code):

    if language == "Solidity":
        pattern = r"(?s)pragma solidity.*\}"
        matches = re.findall(pattern, code)
        solidity_code = ''.join(matches)        
        filepath = './TestFiles'
        split = filename.split(".")
        file_base_name = split[0]+"_test"
        file_extension = '.sol'


    elif language == "Go":
        pattern = r"(?s)package main.*\}"
        matches = re.findall(pattern, code)
        solidity_code = ''.join(matches) 
        filepath = './TestFiles'
        split = filename.split(".")
        file_base_name = split[0]+"_test"
        file_extension = '.go'
    
    else:
        print(f"Error: Unsupported language '{language}'")
        return None
    i=1

    filename = os.path.join(filepath, file_base_name + file_extension)
    while os.path.exists(filename):
        filename = os.path.join(filepath, f'{file_base_name}_{i}{file_extension}')
        i += 1

    try:
        with open(filename, 'w') as f:
            f.write(solidity_code)
            f.close()
    except IOError:
        print(f"Error: Could not write to file '{filename}'")
        return None

    print("Test File Created")

    with open('FileHistory.txt', 'a') as f:
        f.write(filename)
        f.write("\n")
        f.close()
    split = filename.split("/")
    filename = split[2]
    return filename