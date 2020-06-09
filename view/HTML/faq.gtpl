{{ define "faq" }}
<!DOCTYPE html>
<html lang="en">
  <head>
    {{template "head"}}
  </head>
  <body>
    {{template "navbar"}}
    <!-- Page Content -->
    <div class="container">
      <!-- FAQ Accordion -->
      <div class="row">
        <div class="panel-group" id="accordion" role="tablist" aria-multiselectable="true">
          {{ range $idx, $element := . }}
          <div class="panel panel-default">
            <div class="panel-heading" role="tab" id="heading{{ $idx }}">
              <h4 class="panel-title">
                <a {{ if gt $idx 0 }}class="collapsed"{{ end }} role="button" data-toggle="collapse" data-parent="#accordion" href="#collapse{{ $idx }}" aria-expanded={{ if gt $idx 0 }}"false"{{ else }}"true"{{ end }} aria-controls="collapse{{ $idx }}">
                  {{ index $element 0 }}
                </a>
              </h4>
            </div>
            <div id="collapse{{ $idx }}" class="panel-collapse collapse {{ if eq $idx 0 }}in{{ end }}" role="tabpanel" aria-labelledby="heading{{ $idx }}">
              <div class="panel-body">
                {{ if gt $idx 0 }}
                  {{ index $element 1 }}
                {{ else }}
                  <div class="list-group">
                  {{ $curTitle := "" }}
                  {{ range $i, $e := $element }}
                    {{ if gt $i 0 }}
                      {{ if IsEven $i }}
                        <a href="{{ $e }}" class="list-group-item list-group-item-info">{{ $curTitle }}</a>
                      {{ else }}
                        {{ $curTitle = $e }}
                      {{ end }}
                    {{ end }}
                  {{ end }}
                </div>
                {{ end }}
              </div>
            </div>
          </div>
          {{ end }}
        </div>
      </div><!-- /FAQ Accordion row -->
    </div>
  <!-- /.container -->

    {{template "jsfoot"}}
  </body>
</html>
{{ end }}
