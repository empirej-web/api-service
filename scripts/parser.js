const fs = require('fs');
const path = require('path');

class Parser {
  constructor(filePath) {
    this.filePath = filePath;
  }

  async parse() {
    try {
      const fileContent = await fs.promises.readFile(this.filePath, 'utf8');
      const jsonData = JSON.parse(fileContent);
      return jsonData;
    } catch (error) {
      if (error.code === 'ENOENT') {
        throw new Error(`File not found: ${this.filePath}`);
      } else if (error instanceof SyntaxError) {
        throw new Error(`Invalid JSON in file: ${this.filePath}`);
      } else {
        throw error;
      }
    }
  }

  async parseCsv() {
    try {
      const fileContent = await fs.promises.readFile(this.filePath, 'utf8');
      const rows = fileContent.split('\n');
      const headers = rows.shift().split(',');
      const data = rows.map(row => {
        const values = row.split(',');
        return headers.reduce((obj, header, index) => {
          obj[header.trim()] = values[index].trim();
          return obj;
        }, {});
      });
      return data;
    } catch (error) {
      if (error.code === 'ENOENT') {
        throw new Error(`File not found: ${this.filePath}`);
      } else {
        throw error;
      }
    }
  }

  async parseXml() {
    try {
      const fileContent = await fs.promises.readFile(this.filePath, 'utf8');
      const parser = new DOMParser();
      const xmlDoc = parser.parseFromString(fileContent, 'text/xml');
      const data = [];
      const elements = xmlDoc.getElementsByTagName('*');
      for (let i = 0; i < elements.length; i++) {
        const element = elements[i];
        const obj = {};
        for (let j = 0; j < element.attributes.length; j++) {
          const attribute = element.attributes[j];
          obj[attribute.nodeName] = attribute.nodeValue;
        }
        data.push(obj);
      }
      return data;
    } catch (error) {
      if (error.code === 'ENOENT') {
        throw new Error(`File not found: ${this.filePath}`);
      } else if (error instanceof Error && error.message.includes('parsererror')) {
        throw new Error(`Invalid XML in file: ${this.filePath}`);
      } else {
        throw error;
      }
    }
  }
}

module.exports = Parser;