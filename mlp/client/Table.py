from typing import List


class Table:
    cols: List[str]
    rows: List[List[str]]

    def __init__(self, *args):
        self.cols = args
        self.rows = []
        self.styles =[]
        for i in range(len(args)):
            self.styles.append([])

        self.styles[0] = ['text-align:left']

    def row(self, *args):
        self.rows.append(args)

    def monospace(self, pos=0):
        self.styles[pos].append('font-family: monospace;')

    def _repr_html_(self):

        head = ""
        for text, styles in zip(self.cols, self.styles):
            style = ''
            if styles:
                style = ' style="' + ';'.join(styles) + '"'
            head += f'<th {style}>{text}</th>'

        body = ''

        for row in self.rows:
            body += '<tr>'

            for text, styles in zip(row, self.styles):
                style = ''
                if styles:
                    style = ' style="' + ';'.join(styles) + '"'
                body += f'<td {style}>{text}</td>'

            body += '</tr>\n'

        s = f"""
    
<table class='table table-striped' style='table-layout: auto;'>
    <thead>
        <tr>
            {head}
        </tr>
    </thead>
    <tbody>
        {body}
    </tbody>
</table>"""
        return s
